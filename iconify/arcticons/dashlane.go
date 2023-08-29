package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dashlane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.483 5.976l-4.738-1.432a1.025 1.025 0 0 0-1.322.982v35.517c0 .452.296.85.729.981l4.738 1.431a1.025 1.025 0 0 0 1.322-.981V6.957c0-.452-.297-.85-.73-.981Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.271 5.976l-4.738-1.432a1.025 1.025 0 0 0-1.321.982v9.787c0 .452.296.85.728.982l4.738 1.43a1.025 1.025 0 0 0 1.322-.98V6.956c0-.452-.296-.85-.729-.981Zm0 25.729l-4.738-1.43a1.025 1.025 0 0 0-1.321.98v9.788c0 .452.296.85.728.981l4.738 1.431A1.025 1.025 0 0 0 24 42.474v-9.787c0-.452-.296-.851-.729-.982ZM30.06 9.75l-4.738-1.43A1.025 1.025 0 0 0 24 9.3v9.788c0 .452.296.85.729.981l4.738 1.431a1.025 1.025 0 0 0 1.321-.981v-9.787c0-.452-.296-.851-.728-.982Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.06 29.63l-4.738-1.432A1.025 1.025 0 0 0 24 29.18v9.787c0 .452.296.85.729.981l4.738 1.431a1.025 1.025 0 0 0 1.321-.981V30.61c0-.452-.296-.851-.728-.982Zm6.788-13.885l-4.738-1.43a1.025 1.025 0 0 0-1.322.98V33.84c0 .452.296.85.73.981l4.737 1.431a1.025 1.025 0 0 0 1.322-.981V16.727c0-.452-.296-.851-.729-.982Z"/>`),
		g.Group(children),
	)
}