package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinerWorkerAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M33.655 20.75h-8.172v19.615A1.64 1.64 0 0 1 23.848 42a1.64 1.64 0 0 1-1.635-1.635v-8.172h-3.269v8.172A1.64 1.64 0 0 1 17.31 42a1.64 1.64 0 0 1-1.635-1.635V20.75v7.467c-.899 0-3.675-1.711-3.675-5.925c0-3.176 2.776-4.81 3.675-4.81h17.98a1.64 1.64 0 0 1 1.635 1.634a1.64 1.64 0 0 1-1.635 1.634Z"/><path d="M34 11v30h-2V11h2Z"/><path d="M26 15c0-1.657 3.134-3 7-3s7 1.343 7 3c-5-1.5-8.5-1.5-14 0Z"/><path fill-rule="evenodd" d="M25 10.5h-1.902c.468.566.75 1.29.75 2.078a3.279 3.279 0 0 1-3.27 3.269a3.279 3.279 0 0 1-3.268-3.27c0-.787.282-1.511.75-2.077H16a4.5 4.5 0 1 1 9 0Zm-4.5-.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}