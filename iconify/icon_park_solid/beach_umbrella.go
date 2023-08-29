package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeachUmbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBeachUmbrella0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M14.34 21.412L24 24l9.66 2.588l9.659 2.588c2.338-8.725-1.471-17.65-8.773-22.176a20 20 0 0 0-5.37-2.318A19.97 19.97 0 0 0 22.5 4.05c-8.237.62-15.56 6.342-17.818 14.773l9.659 2.588Z"/><path stroke="#000" d="M29.176 4.681s-5.64 3.665-8.866 7.977s-4.037 9.271-4.037 9.271M29.177 4.681s3.052 5.995 3.69 11.342c.638 5.347-1.14 10.047-1.14 10.047"/><path stroke="#fff" d="m4.682 18.824l9.659 2.588L24 24l9.658 2.588l9.66 2.588M4 44h40M24 24l-5.5 20m4-39.949a19.97 19.97 0 0 1 6.676.63a20 20 0 0 1 5.37 2.32"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBeachUmbrella0)"/>`),
		g.Group(children),
	)
}