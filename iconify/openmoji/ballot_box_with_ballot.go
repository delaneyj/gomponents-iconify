package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallotBoxWithBallot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" stroke="#d0cfce" stroke-linejoin="round" stroke-width="2" d="M10.14 63.05V31.1h51.73v31.95z"/><path fill="#9b9b9a" stroke="#9b9b9a" stroke-linecap="round" stroke-width="2" d="M61.35 37L35.01 63.05h10.23L61.35 37"/><path fill="#d0cfce" d="M22.17 36.99h27.87v3.752H22.17z"/><path fill="#fff" d="m24.1 31.34l16.17 5.836l9.927-22.72l-16.38-6.695z"/><path fill="#9b9b9a" stroke="#9b9b9a" stroke-linejoin="round" stroke-width="2" d="M45.24 63.05h16.62L61.345 37z"/><path fill="none" stroke="#9b9b9a" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21.97 38.89h28.18"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M43.706 31.096h18.15v29.942m-51.722 0V31.096h13.534m13.595 5.745l-13.486-5.529l9.988-24.364l16.72 6.851l-9.376 22.862m-19.199.259H50.1"/>`),
		g.Group(children),
	)
}