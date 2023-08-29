package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="16" cy="25" r="2" fill="currentColor"/><path fill="currentColor" d="m10.47 19.233l1.414 1.413a5.969 5.969 0 0 1 8.223-.009l1.413-1.413a7.963 7.963 0 0 0-11.05.01Z"/><path fill="currentColor" d="m6.229 14.993l1.414 1.413a11.955 11.955 0 0 1 16.704-.01l1.413-1.413a13.951 13.951 0 0 0-19.531.01Z"/><path fill="currentColor" d="M30 10.741a19.94 19.94 0 0 0-28 0v.023l1.404 1.404a17.934 17.934 0 0 1 25.181-.01L30 10.743Z"/>`),
		g.Group(children),
	)
}