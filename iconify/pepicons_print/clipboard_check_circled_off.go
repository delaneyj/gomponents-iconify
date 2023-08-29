package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCheckCircledOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path fill-rule="evenodd" d="M8 4.351c0-.47.414-.851.926-.851h6.148c.512 0 .926.381.926.851V7.65c0 .47-.414.851-.926.851H8.926C8.414 8.5 8 8.119 8 7.649V4.35Z" clip-rule="evenodd"/><path d="M6.462 19h10.576c.532 0 .962-.448.962-1V6c0-.552-.43-1-.962-1H6.462C5.93 5 5.5 5.448 5.5 6v12c0 .552.43 1 .962 1Z"/><path d="M20 15.75a4.25 4.25 0 1 1-8.5 0a4.25 4.25 0 0 1 8.5 0Z"/></g><path fill-rule="evenodd" d="M6.175 2.5a.5.5 0 0 1 .5-.5h6.643a.5.5 0 0 1 .5.5v3.875a.5.5 0 0 1-.5.5H6.675a.5.5 0 0 1-.5-.5V2.5Zm1 .5v2.875h5.643V3H7.175Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M4.5 17V5h2V4h-2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h7.854a4.02 4.02 0 0 1-.819-1H4.5Zm11-5.97c.35.045.685.133 1 .26V5a1 1 0 0 0-1-1h-2v1h2v6.03Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15 18a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 1a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M16.582 13.882a.5.5 0 0 1 .078.703l-1.106 1.382a1 1 0 0 1-1.488.082l-.696-.695a.5.5 0 0 1 .708-.707l.696.695l1.105-1.382a.5.5 0 0 1 .703-.078Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}