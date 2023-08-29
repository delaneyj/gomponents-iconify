package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperboardCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilClapperboardCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M3.768 7.972a.5.5 0 0 1 .382-.595l14.653-3.208a.5.5 0 0 1 .595.381l.642 2.93a.5.5 0 0 1-.382.596L5.005 11.284a.5.5 0 0 1-.595-.381l-.642-2.93Zm1.084.275L5.28 10.2l13.676-2.995l-.428-1.954L4.852 8.247Z"/><path d="M12.854 9.338L10.26 6.835l.694-.72l2.596 2.503l-.695.72Zm-3.907.855L6.352 7.691l.694-.72L9.64 9.474l-.694.72Zm7.815-1.711L14.166 5.98l.695-.72l2.595 2.503l-.694.72Zm-4.773 5.795l2-3l-.832-.554l-2 3l.832.554Zm4 0l2-3l-.832-.554l-2 3l.832.554Zm-8 0l2-3l-.832-.554l-2 3l.832.554Z"/><path d="M4.573 11a.5.5 0 0 1 .5-.5h15a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-.5.5h-15a.5.5 0 0 1-.5-.5v-9Zm1 .5v8h14v-8h-14Z"/><path d="M20.573 14.5h-16v-1h16v1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilClapperboardCircleFilled0)"/></g>`),
		g.Group(children),
	)
}