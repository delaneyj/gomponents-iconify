package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webinsta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1011 298L785 715q-18 33-53.5 44T641 770q-27 0-49.5-11T557 732L395 440q-16-28-7-59.5t37.5-48.5t60-9t47.5 36l96 174l189-350q23-39 67-50.5t83.5 12t52 68.5t-9.5 85zm-498-41q-53 0-90.5-37.5t-37.5-91t37.5-91T513 0t90.5 37.5t37.5 91t-37.5 91T513 257zM391 754.5q-40 23.5-83.5 12T241 715L15 298q-22-40-9.5-85t52-68.5t83.5-12t67 50.5l226 418q22 40 9.5 85T391 754.5z"/>`),
		g.Group(children),
	)
}