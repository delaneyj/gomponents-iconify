package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrownCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilCrownCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="m17.896 16.818l1.515-5.766l-2.214 1.41a2 2 0 0 1-2.74-.578L13 9.695l-1.458 2.19a2 2 0 0 1-2.74.577l-2.213-1.41l1.515 5.766h9.792Zm-10.77-6.61c-.767-.489-1.736.218-1.505 1.098l1.516 5.766a1 1 0 0 0 .967.746h9.792a1 1 0 0 0 .967-.746l1.516-5.766c.23-.88-.738-1.586-1.505-1.098l-2.214 1.41a1 1 0 0 1-1.37-.288l-1.458-2.19a1 1 0 0 0-1.664 0l-1.458 2.19a1 1 0 0 1-1.37.289l-2.214-1.41Z" clip-rule="evenodd"/><path d="M13.944 6.945a.945.945 0 1 1-1.89.002a.945.945 0 0 1 1.89-.002ZM21.5 8.836a.945.945 0 1 1-1.89.001a.945.945 0 0 1 1.89 0Zm-15.111 0a.945.945 0 1 1-1.89.001a.945.945 0 0 1 1.89 0Z"/><path fill-rule="evenodd" d="M8.25 19a.5.5 0 0 1 .5-.5h8.737a.5.5 0 1 1 0 1H8.75a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilCrownCircleFilled0)"/></g>`),
		g.Group(children),
	)
}