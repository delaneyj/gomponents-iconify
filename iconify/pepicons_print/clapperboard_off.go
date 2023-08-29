package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperboardOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="m3.156 6.137l13.816-3.025l.606 2.764L3.76 8.9l-.605-2.764Z"/><path fill-rule="evenodd" d="M2.18 6.351a1 1 0 0 1 .762-1.19L16.76 2.135a1 1 0 0 1 1.19.763l.605 2.764a1 1 0 0 1-.763 1.19L3.975 9.877a1 1 0 0 1-1.19-.762L2.178 6.35Zm2.167.55l.177.809l11.863-2.597l-.177-.81L4.347 6.9Z" clip-rule="evenodd"/><path d="M3.926 9.093h14.143v8.486H3.926V9.093Z"/><path fill-rule="evenodd" d="M2.926 9.093a1 1 0 0 1 1-1h14.143a1 1 0 0 1 1 1v8.486a1 1 0 0 1-1 1H3.926a1 1 0 0 1-1-1V9.093Zm2 1v6.486h12.143v-6.486H4.926Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M.768 4.972a.5.5 0 0 1 .382-.595l14.652-3.208a.5.5 0 0 1 .596.381l.642 2.93a.5.5 0 0 1-.382.596L2.005 8.284a.5.5 0 0 1-.595-.381l-.642-2.93Zm1.084.275L2.28 7.2l13.676-2.995l-.428-1.954L1.852 5.247Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.854 6.338L7.26 3.835l.694-.72l2.596 2.503l-.695.72Zm-3.907.855L3.352 4.691l.694-.72L6.64 6.474l-.694.72Zm7.815-1.711L11.166 2.98l.695-.72l2.595 2.503l-.694.72Zm-4.773 5.795l2-3l-.832-.554l-2 3l.832.554Zm4 0l2-3l-.832-.554l-2 3l.832.554Zm-8 0l2-3l-.832-.554l-2 3l.832.554Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M1.573 8a.5.5 0 0 1 .5-.5h15a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-.5.5h-15a.5.5 0 0 1-.5-.5V8Zm1 .5v8h14v-8h-14Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17.573 11.5h-16v-1h16v1Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}