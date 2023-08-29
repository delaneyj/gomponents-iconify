package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g transform="translate(3 3)"><path fill-rule="evenodd" d="m6.804 2.632l-.637.264A3.507 3.507 0 0 0 4 6.137v4.386a1.46 1.46 0 0 0 2.167 1.276l.227-.126a4 4 0 0 1 3.88 0l.453.251a4 4 0 0 0 3.88 0l.734-.407A3.222 3.222 0 0 0 17 8.7V4.638a1.605 1.605 0 0 0-2.07-1.534l-.893.272a4 4 0 0 1-2.694-.131l-1.48-.613a4 4 0 0 0-3.059 0Zm4.893 7.543l-.454-.251A6 6 0 0 0 6 9.644V6.136c0-.61.368-1.16.931-1.393l.638-.263a2 2 0 0 1 1.529 0l1.481.612a6 6 0 0 0 4.04.196L15 5.173V8.7c0 .444-.241.853-.63 1.068l-.733.407a2 2 0 0 1-1.94 0Z" clip-rule="evenodd"/><rect width="2" height="16" x="4" y="2" rx="1"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}