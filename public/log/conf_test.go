/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:569929309

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:569929309

	2017.7 by hunterhug
*/
package log

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestLoadConf(t *testing.T) {
	convey.Convey("测试配置", t, func() {
		var js = `
		{
			"Appenders":{
				"test_appender":{
					"Type":"file",
					"Target":"/tmp/test.log"
				},
				"a_appender":{
					"Type":"console"
				},
				"b_appender":{
					"Type":"file",
					"Target":"/tmp/test.b.log"
				}
			},
			"Loggers":{
				"sunteng/commons/log/a":{
					"Appenders":["test_appender","a_appender"],
					"Level":"debug"
				},
				"sunteng/commons/log/b":{
					"Level":"error"
				}
			},
			"Root":{
				"Level":"log",
				"Appenders":["test_appender"]
			}
		}
		`
		_, err := LoadConf(js)
		convey.So(err, convey.ShouldEqual, nil)
	})
}
