package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FramedPicture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M49 34.446V17H15v30h34V34.446zM41.999 20a4 4 0 1 1 .003 8a4 4 0 0 1-.003-8zm-9.171 9.683l8.653 8.064l-8.653 2.606v-10.67zm14.66 15.985H16.512V34.446h7.614l-4.177 3.893l12.124 3.652l12.124-3.652l-4.178-3.893h7.469v11.222zM31.317 29.683v10.671l-8.652-2.606l8.652-8.065z"/><path fill="currentColor" d="M2 4v56h60V4H2zm51 47H11V13h42v38z"/>`),
		g.Group(children),
	)
}