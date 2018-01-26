func Test(ctrip string)(m map[string]domain.Test,err error){
	headData := genBaseQCityData()
  
  ///////////////////////request
	q := u.Query()
	q.Set("reqData",string(body))
	u.RawQuery = q.Encode()

	httpClient := &httputils.ServiceHttpClient{}
	response := staticdata_api.QCityResponse{}
	url := ""
  code, message, data := httpClient.GetJson(url ,map[string]string{"Connection":"keep-alive"}, 1, nil)//todo url
	if code != commobj.SUCCESS{
		log.Println("MESSAGE====",message)
		return
	}
	enc := mahonia.NewDecoder("gbk")
	temp := enc.ConvertString(string(data))
	err = json.Unmarshal([]byte(temp),&response)
	if err != nil{
		log.Println(err)
	}
	xlsx := excelize.NewFile()
	categories := map[string]string{"A1": "CountryId","B1": "Country", "C1": "ProvinceId", "D1": "Province","E1": "CityId","F1":"City"}
	for k, v := range categories {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	for index,info := range response.CityResult.CityList {
		xlsx.SetCellValue("sheet1","A"+strconv.Itoa(index+2),info.CountryId)
		xlsx.SetCellValue("sheet1","B"+strconv.Itoa(index+2),info.Country)
		xlsx.SetCellValue("sheet1","C"+strconv.Itoa(index+2),info.ProvinceId)
		xlsx.SetCellValue("sheet1","D"+strconv.Itoa(index+2),info.Province)
		xlsx.SetCellValue("sheet1","E"+strconv.Itoa(index+2),info.CityId)
		xlsx.SetCellValue("sheet1","F"+strconv.Itoa(index+2),info.City)
		fmt.Println(index)
	}
	log.Println("开始保存")
	err = xlsx.SaveAs("./Workbook.xlsx")
	log.Println("保存完成")
	if err != nil {
		fmt.Println(err)
	}
	return
}
