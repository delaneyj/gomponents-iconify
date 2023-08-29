package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilCarCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.298 6.954L6.1 11.651a2 2 0 0 0-1.601 1.96V16.5a2 2 0 0 0 2 2h13a2 2 0 0 0 2-2v-2.889a2 2 0 0 0-1.601-1.96l-1.197-4.697a2 2 0 0 0-1.475-1.452A18.211 18.211 0 0 0 13 5c-1.41 0-2.82.168-4.228.502a2 2 0 0 0-1.475 1.452ZM6.5 12.611a.5.5 0 0 0 .485-.376L8.267 7.2a1 1 0 0 1 .738-.726A17.224 17.224 0 0 1 13 6c1.33 0 2.662.158 3.995.475a1 1 0 0 1 .737.726l1.282 5.034a.5.5 0 0 0 .485.376a1 1 0 0 1 1 1V16.5a1 1 0 0 1-1 1h-13a1 1 0 0 1-1-1v-2.889a1 1 0 0 1 1-1Z"/><path d="M8.5 19a1.5 1.5 0 0 0-3 0v1a1.5 1.5 0 0 0 3 0v-1Zm-2 0a.5.5 0 0 1 1 0v1a.5.5 0 0 1-1 0v-1Zm14 0a1.5 1.5 0 0 0-3 0v1a1.5 1.5 0 0 0 3 0v-1Zm-2 0a.5.5 0 0 1 1 0v1a.5.5 0 0 1-1 0v-1Zm-.5-2.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0-2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1Zm-10 2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0-2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1Zm7.648-7h-5.296a1.5 1.5 0 0 0-1.457 1.143l-.49 2A1.5 1.5 0 0 0 9.862 12.5h6.276a1.5 1.5 0 0 0 1.457-1.857l-.49-2A1.5 1.5 0 0 0 15.647 7.5ZM9.866 8.881a.5.5 0 0 1 .486-.381h5.296a.5.5 0 0 1 .485.381l.49 2a.5.5 0 0 1-.485.619H9.862a.5.5 0 0 1-.486-.619l.49-2Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilCarCircleFilled0)"/></g>`),
		g.Group(children),
	)
}