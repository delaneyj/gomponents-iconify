package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Serverless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0-.002v106.617h169.305L197.38-.002H0zm286.7 0l-28.077 106.617H512V-.002H286.7zM0 202.693v106.612h115.926l28.076-106.612H0zm233.322 0l-28.076 106.612H512V202.693H233.322zM0 405.387v106.611h62.545l28.08-106.611H0zm179.941 0l-28.076 106.611h360.133V405.387H179.941z"/>`),
		g.Group(children),
	)
}