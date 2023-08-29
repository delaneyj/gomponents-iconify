package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bicycle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M13.062 3.936a1 1 0 1 1-.124-1.996C16.828 1.7 19 2.46 19 4.44c0 1.574-.978 2.274-2.855 2.55a1 1 0 1 1-.29-1.98C16.87 4.862 17 4.77 17 4.44c0-.296-1.094-.68-3.938-.504Z"/><path fill-rule="evenodd" d="M15.5 20a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm0-7a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Zm-11 7a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm0-7a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Z" clip-rule="evenodd"/><path d="M5 6.5a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2H5Z"/><path fill-rule="evenodd" d="M4.5 16.5H10a1 1 0 0 0 .983-1.185l3.864-6.165A1 1 0 0 0 14 7.619H8.45l-.802-2.432a1 1 0 0 0-1.9.627L6.755 8.86l-3.146 6.186A1 1 0 0 0 4.5 16.5Zm3.138-4.964L6.13 14.5h2.487l-.979-2.964Zm2.536 1.303l2.019-3.22H9.111l1.063 3.22Z" clip-rule="evenodd"/><path d="m13.98 2.743l2.5 12.562c.261 1.308-1.7 1.698-1.96.39l-2.5-12.562c-.261-1.307 1.7-1.698 1.96-.39Z"/></g>`),
		g.Group(children),
	)
}