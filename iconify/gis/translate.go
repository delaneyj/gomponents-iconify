package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Translate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="m65.215 39.958l-9.838 6.318L45.32 30.613l-10.2 6.55l10.059 15.663l-9.837 6.317l24.53 5.345z" color="currentColor"/><path fill="currentColor" fill-rule="evenodd" d="M69.475 2a3.5 3.5 0 0 0-2.659 1.14l-9.822 10.758l-34.035 2.09a3.5 3.5 0 0 0-2.14.905L1.144 34.787A3.5 3.5 0 0 0 .91 39.73a3.5 3.5 0 0 0 4.945.235L24.617 22.9l34.22-2.101a3.5 3.5 0 0 0 2.37-1.133l10.78-11.807a3.5 3.5 0 0 0-.225-4.943A3.5 3.5 0 0 0 69.475 2Zm26.834 58.94l-3.774 4.132l3.692 3.371L100 64.311Zm-7.145 7.824l-2.951 3.232l-3.592.22l.307 4.991l4.601-.281a2.5 2.5 0 0 0 1.694-.81l3.632-3.981zm-11.533 3.76l-9.983.613l.307 4.99l9.983-.613zm-14.973.92l-9.98.61l.306 4.993l9.98-.613zm-15.443 4.068l-7.397 6.728l3.364 3.7l7.398-6.731zM36.119 87.604L30.57 92.65l3.366 3.7l5.546-5.047z" color="currentColor"/>`),
		g.Group(children),
	)
}