package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Faceapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.366 39.572a22.03 22.03 0 0 0 1.436-7.934c0-8.084-5.296-11.199-10.724-13.12c-5.593-1.979-11.322-2.711-11.322-7.728c0-4.633 4.113-6.29 8.986-6.29c5.152 0 18.212 3.654 18.212 18.751a23.752 23.752 0 0 1-6.588 16.32"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.253 14.405c0 3.934-5.132 7.888-5.132 11.362a11.805 11.805 0 0 0 1.877 6.15s-4.952-3.394-4.952-11.022S11.38 9.625 14.145 8.55m8.557 34.95c2.804 0 12.718-5.236 12.718-15.336"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.702 43.5c-2.624 0-11.475-4.586-12.6-13.453"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.348 34.35a4.11 4.11 0 0 1-3.583-4.469a3.81 3.81 0 0 1 .75-2.3"/>`),
		g.Group(children),
	)
}