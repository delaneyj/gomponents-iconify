package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaxiCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilTaxiCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M9.171 6a2 2 0 0 0-1.94 1.515L6.1 12.04A2 2 0 0 0 4.5 14v3a2 2 0 0 0 2 2h13a2 2 0 0 0 2-2v-3a2 2 0 0 0-1.6-1.96l-1.13-4.525A2 2 0 0 0 16.828 6H9.17ZM6.5 13a.5.5 0 0 0 .485-.379l1.216-4.864A1 1 0 0 1 9.171 7h7.658a1 1 0 0 1 .97.757l1.216 4.864A.5.5 0 0 0 19.5 13a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1h-13a1 1 0 0 1-1-1v-3a1 1 0 0 1 1-1Z"/><path d="M18 17a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0-2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1ZM8 17a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0-2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1Zm.5 4.5a1.5 1.5 0 0 0-3 0v1a1.5 1.5 0 0 0 3 0v-1Zm-2 0a.5.5 0 0 1 1 0v1a.5.5 0 0 1-1 0v-1Zm14 0a1.5 1.5 0 0 0-3 0v1a1.5 1.5 0 0 0 3 0v-1Zm-2 0a.5.5 0 0 1 1 0v1a.5.5 0 0 1-1 0v-1ZM15.648 8h-5.296a1.5 1.5 0 0 0-1.457 1.143l-.49 2A1.5 1.5 0 0 0 9.862 13h6.276a1.5 1.5 0 0 0 1.457-1.857l-.49-2A1.5 1.5 0 0 0 15.647 8ZM9.866 9.381A.5.5 0 0 1 10.352 9h5.296a.5.5 0 0 1 .485.381l.49 2a.5.5 0 0 1-.485.619H9.862a.5.5 0 0 1-.486-.619l.49-2ZM14.5 5.5a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Zm-2 0a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilTaxiCircleFilled0)"/></g>`),
		g.Group(children),
	)
}