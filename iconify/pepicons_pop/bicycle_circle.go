package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BicycleCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M16.062 6.936a1 1 0 1 1-.124-1.996C19.828 4.7 22 5.46 22 7.44c0 1.574-.978 2.274-2.855 2.55a1 1 0 1 1-.29-1.98C19.87 7.862 20 7.77 20 7.44c0-.296-1.094-.68-3.938-.504Z"/><path fill-rule="evenodd" d="M18.5 23a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm0-7a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Zm-11 7a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm0-7a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Z" clip-rule="evenodd"/><path d="M8 9.5a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2H8Z"/><path fill-rule="evenodd" d="M7.5 19.5H13a1 1 0 0 0 .983-1.185l3.864-6.165A1 1 0 0 0 17 10.619h-5.55l-.802-2.432a1 1 0 0 0-1.9.627l1.007 3.047l-3.146 6.186A1 1 0 0 0 7.5 19.5Zm3.138-4.964L9.13 17.5h2.487l-.979-2.964Zm2.536 1.303l2.019-3.22h-3.082l1.063 3.22Z" clip-rule="evenodd"/><path d="m16.98 5.743l2.5 12.562c.261 1.308-1.7 1.698-1.96.39l-2.5-12.562c-.261-1.307 1.7-1.698 1.96-.39Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}