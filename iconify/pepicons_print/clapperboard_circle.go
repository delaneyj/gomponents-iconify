package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperboardCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g opacity=".2"><path d="m6.156 9.137l13.816-3.025l.606 2.764L6.76 11.9l-.605-2.764Z"/><path fill-rule="evenodd" d="M5.18 9.351a1 1 0 0 1 .762-1.19L19.76 5.135a1 1 0 0 1 1.19.763l.605 2.764a1 1 0 0 1-.763 1.19L6.975 12.877a1 1 0 0 1-1.19-.762L5.178 9.35Zm2.167.55l.177.809l11.863-2.597l-.177-.81L7.347 9.9Z" clip-rule="evenodd"/><path d="M6.926 12.093h14.143v8.486H6.926v-8.486Z"/><path fill-rule="evenodd" d="M5.926 12.093a1 1 0 0 1 1-1h14.143a1 1 0 0 1 1 1v8.486a1 1 0 0 1-1 1H6.926a1 1 0 0 1-1-1v-8.486Zm2 1v6.486h12.143v-6.486H7.926Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M3.768 7.972a.5.5 0 0 1 .382-.595l14.653-3.208a.5.5 0 0 1 .595.381l.642 2.93a.5.5 0 0 1-.382.596L5.005 11.284a.5.5 0 0 1-.595-.381l-.642-2.93Zm1.084.275L5.28 10.2l13.676-2.995l-.428-1.954L4.852 8.247Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.854 9.338L10.26 6.835l.694-.72l2.596 2.503l-.695.72Zm-3.907.855L6.352 7.691l.694-.72L9.64 9.474l-.694.72Zm7.815-1.711L14.166 5.98l.695-.72l2.595 2.503l-.694.72Zm-4.773 5.795l2-3l-.832-.554l-2 3l.832.554Zm4 0l2-3l-.832-.554l-2 3l.832.554Zm-8 0l2-3l-.832-.554l-2 3l.832.554Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M4.573 11a.5.5 0 0 1 .5-.5h15a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-.5.5h-15a.5.5 0 0 1-.5-.5v-9Zm1 .5v8h14v-8h-14Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20.573 14.5h-16v-1h16v1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}