#! /usr/bin/env python
# coding:utf-8

from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.action_chains import ActionChains
import time
import urllib2
import sys
import re
import csv
import numpy as np

# 解决中文报错的问题
reload(sys)
sys.setdefaultencoding('utf-8')
# 打开一个火狐浏览器
driver = webdriver.Firefox()
# 睡眠3秒，防止浏览器还没打开就进行了其他操作
time.sleep(3)
# 化工商户页面的url
url = 'https://s.1688.com/company/company_search.htm?keywords=%CA%B3%C6%B7&button_click=top&earseDirect=false&n=y&_source=sug'
# 登录的url
login_url = 'https://login.taobao.com/member/login.jhtml'
# 跳转到登录页面
driver.get(login_url)
# 睡眠5秒，防止网速较差打不开网页就进行了其他操作
time.sleep(5)
# # 找到账号登录框的DOM节点，并且在该节点内输入账号
# driver.find_element_by_name("TPL_username").send_keys('15015287866')
# # 找到账号密码框的DOM节点，并且在该节点内输入密码
# driver.find_element_by_name("TPL_password").send_keys('ABCabc123')
# # 找到账号登录框的提交按钮，并且点击提交
# driver.find_element_by_name("TPL_password").send_keys(Keys.ENTER)
# 睡眠5秒，防止未登录就进行了其他操作
time.sleep(9)
# 跳转到化工商户页面的url
driver.get(url)
# 新建一个data.csv文件，并且将数据保存到csv中
csvfile = file('data.csv', 'w')
writer = csv.writer(csvfile)
# 写入标题，我们采集企业名称，主页，产品，联系人，电话和地址信息
writer.writerow((
    u'企业名称'.encode('gbk'),
    u'联系人'.encode('gbk'),
    u'电话'.encode('gbk'),

))
# 构建agents防止反爬虫
user_agents = [
    'User-Agent:Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50',
    'User-Agent:Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0',
    'User-Agent:Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729; InfoPath.3; rv:11.0) like Gecko',
    'Mozilla/5.0 (compatible; Konqueror/3.5; Linux) KHTML/3.5.5(like Gecko) (Kubuntu)',
    'User-Agent": "Mozilla/5.0 (Linux; Android 4.4.4; HTC D820u Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.89 Mobile Safari/537.36',
    'User-Agent: Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)',
    "Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Ubuntu/11.04 Chromium/16.0.912.77 Chrome/16.0.912.77 Safari/535.7",
    "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:10.0) Gecko/20100101 Firefox/10.0 ",
]
# 总共有100页，使用for循环采集
for page in xrange(1, 100):
    # 捕捉异常
    try:
        # 获取企业名称列表
        title = driver.find_elements_by_css_selector("a[class=list-item-title-text]")
        # 获取产品
        product = driver.find_elements_by_xpath("//div[@class=\"list-item-detail\"]/div[1]/div[1]/a[1]")
        # 打印长度，调试
        print len(title)
        # 定义电话正则
        tel_pattern = re.compile('<dd cla.*?>.*?(\w{11})\s.*?</dd>', re.S)
        # 定义移动电话正则
        member_name_pattern = re.compile('<a.*?class="membername".*?>(.*?)</a>', re.S)
        # 定义地址正则
        address_pattern = re.compile('"address">(.*?)</dd>', re.S)
        for i in xrange(len(title)):
            # 获取标题的值
            title_value = title[i].get_attribute('title')
            # 获取跳转的url
            href_value = title[i].get_attribute('href') + 'page/contactinfo.htm'
            # 获取经营范围
            product_value = product[i].text
            # 随机选择agent进行访问
            agent = np.random.choice(user_agents)
            # 组建header头部
            headers = {'User-Agent': agent, 'Accept': '*/*', 'Referer': 'https://s.1688.com/company/company_search.htm?keywords=%CA%B3%C6%B7&button_click=top&earseDirect=false&n=y&_source=sug'}
            # 使用urllib2进行Request
            request = urllib2.Request(href_value, headers=headers)
            # 访问链接
            print "requestUrl",str(request)
            response = urllib2.urlopen(request)
            # 获得网页源码
            html = response.read()
            # 进行信息匹配
            tel = re.findall(tel_pattern, html)
            print "TEL=========",tel
            try:
                tel = tel[0]
                tel = tel.strip()
                print "tel"
            except Exception, e:
                print "exception"
                continue
            print "正则开始"
            member_name = re.findall(member_name_pattern, html)
            try:
                print "one"
                member_name = member_name[0]
                member_name = member_name.strip()
            except Exception, e:
                print "two"
                continue
            # 打印出信息，方便查看进度
            print 'tel:' + tel
            print 'member_name:' + member_name
            data = (
                title_value.encode('gbk', 'ignore'),
                member_name,
                tel,
            )
            writer.writerow(data)
        js = 'var q=document.documentElement.scrollTop=30000'
        driver.execute_script(js)
        time.sleep(1.5)
        page = driver.find_elements_by_css_selector("a[class=page-next]")
        page = page[0]
        page.click()
        time.sleep(2)
    except Exception, e:
        print 'error',e
        continue
# 关闭csv
csvfile.close()
# 关闭模拟浏览器
driver.close()
