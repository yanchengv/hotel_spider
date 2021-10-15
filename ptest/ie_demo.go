package ptest

import (
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

//使用selenium IE浏览器登录酒店管理系统 并且获取相应的seesionID

func IEdemo() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		seleniumPath    = "selenium-server-standalone-4.0.0-alpha-2.jar"
		geckoDriverPath = "IEDriverServer.exe"
		//geckoDriverPath = "msedgedriver.exe"
		//geckoDriverPath = "chromedriver.exe"
		port = 9038
	)
	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),             // Output debug information to STDERR.

	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	//browserName: chrome MicrosoftEdge internet explorer
	//caps := selenium.Capabilities{"browserName": "chrome"}
	//caps := selenium.Capabilities{"browserName": "MicrosoftEdge"}
	// caps := selenium.Capabilities{"browserName": "internet explorer", "javascriptEnabled": true,
	// 	"acceptSslCerts": true, "allowBlockedContent": true, "ignoreProtectedModeSettings": true, "ignoreZoomSetting": true, "nativeEvents": false, "ignoreZoomLevel": true}
	options := map[string]interface{}{"ignoreZoomSetting": true} //忽略浏览器的大小，IE不加此参数会报错
	caps := selenium.Capabilities{"browserName": "internet explorer", "se:ieOptions": options}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	wd.MaximizeWindow("") //窗口最大化
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("http://121.22.40.202:8083/hotelbs254/Login.aspx"); err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 10)
	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#tbUser")
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	// Enter some new code in text box.
	err = elem.SendKeys("1303250111")
	if err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	elem2, err := wd.FindElement(selenium.ByCSSSelector, "#tbPassword")
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate code already in the text box.
	if err := elem2.Clear(); err != nil {
		panic(err)
	}

	// Enter some new code in text box.
	err = elem2.SendKeys("123456")
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 20)
	// Click the run button.

	btn, err := wd.FindElement(selenium.ByCSSSelector, "#ibtLogin")
	//btn.MoveTo(842, 663)
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}
	//获取sessionId
	s, _ := wd.GetCookie("ASP.NET_SessionId")
	fmt.Println(111111)
	fmt.Println(s.Value)
	fmt.Println(s.Name)
	fmt.Println(2223333333)
	// a, _ := wd.GetCookies() //获取所有cookies
	// fmt.Println(111111111111111111)
	// fmt.Println(a)
	time.Sleep(time.Minute * 5)
}
