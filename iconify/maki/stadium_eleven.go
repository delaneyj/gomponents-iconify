package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StadiumEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5 0v3.012c-2.208.097-3.912.806-4 1.67v3.681C1 9.267 3.014 10 5.5 10S10 9.267 10 8.363V4.682c-.088-.864-1.79-1.574-4-1.67v-.44L8.5 1.5L5 0zM1.818 5.752c.319.178.72.332 1.182.453v2.46c-.75-.237-1.179-.568-1.182-.915V5.752zm7.364.004V7.75c-.002.348-.43.68-1.182.916V6.203c.461-.12.862-.271 1.182-.447zM4 6.398a11.066 11.066 0 0 0 3 0v2.493C6.528 8.962 6.017 9 5.5 9S4.472 8.962 4 8.89V6.399z" fill="currentColor"/>`),
		g.Group(children),
	)
}