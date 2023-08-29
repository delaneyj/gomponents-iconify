package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rdoc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M401.946 421.39v64.969h-64.968l64.968-64.969zm-188.229 31.473H33.496V50.917h234.468l100.487 100.486v1.051h33.495v-34.546L301.46 17.42H0v468.94h247.213l-33.496-33.496zm188.33-280.141H182.144L72.19 282.674l219.904 219.905L512 282.674L402.048 172.722zm-274.88 109.952l82.464-82.464H374.56l82.464 82.464l-164.929 164.929l-164.928-164.929zm247.393-54.976h-82.465v192.416l137.44-137.44l-54.975-54.976z"/>`),
		g.Group(children),
	)
}