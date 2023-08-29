package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileAudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFileAudio0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFileAudio1" fill="currentColor"><path id="feFileAudio2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9.172 2H6v16h12V6.828h-2.828V4ZM15 14.085V11h-4v5.5a1.5 1.5 0 1 1-1-1.415V10a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v5.5a1.5 1.5 0 1 1-1-1.415Z"/></g></g>`),
		g.Group(children),
	)
}