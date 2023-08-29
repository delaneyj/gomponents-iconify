package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VictronconnectAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.425 13.55c-3.05.022-5.16 4.16-5.25 10.7m22.25-1c-.045 7.66-1.87 11.2-5.36 11.2c-5.09-.086-5.5-6.42-5.55-10.2m-1.59-10.7c-3.05.022-5.16 4.16-5.25 10.7"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.825 32.45c-1.58-2.19-1.77-5.79-1.8-8.28m10.9-.92c-.045 7.66-1.87 11.2-5.36 11.2m-7.24-20.9c-3.05.022-5.16 4.16-5.25 10.7m11.25 0c-.049-4.03-.765-6.92-2.07-8.66m12.97 7.66c-.045 7.66-1.87 11.2-5.36 11.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.825 24.25c-.086-7.13-2.27-10.7-6.09-10.7c-3.05.022-5.16 4.16-5.25 10.7m22.34-1c-.045 7.66-1.87 11.2-5.36 11.2"/>`),
		g.Group(children),
	)
}