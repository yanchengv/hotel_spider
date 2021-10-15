package ptest

import (
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

//使用selenium IE浏览器登录酒店管理系统 并且获取相应的seesionID

func Chromedemo() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		seleniumPath = "selenium-server-standalone-4.0.0-alpha-2.jar"
		//geckoDriverPath = "IEDriverServer.exe"
		//geckoDriverPath = "msedgedriver.exe"
		geckoDriverPath = "chromedriver.exe"
		port            = 9037
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
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	wd.MaximizeWindow("") //窗口最大化
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("https://staging.aranya.cc/login"); err != nil {
		panic(err)
	}
	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#phone")
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	// 输入框输入账号.
	err = elem.SendKeys("1111")
	if err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	elem2, err := wd.FindElement(selenium.ByCSSSelector, "#password")
	if err != nil {
		panic(err)
	}
	// Remove the boilerplate code already in the text box.
	if err := elem2.Clear(); err != nil {
		panic(err)
	}

	// 输入框输入密码
	err = elem2.SendKeys("1111")
	if err != nil {
		panic(err)
	}

	// 点击登录.
	btn, err := wd.FindElement(selenium.ByTagName, "button")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}
	s, _ := wd.GetCookie("_aranya_web_session")
	fmt.Println(222222222222222222)
	fmt.Println(s.Value)
	fmt.Println(s.Name)
	fmt.Println(2223333333)
	a, _ := wd.GetCookies()

	fmt.Println(111111111111111111)
	fmt.Println(a)
	time.Sleep(time.Minute * 3)
}
