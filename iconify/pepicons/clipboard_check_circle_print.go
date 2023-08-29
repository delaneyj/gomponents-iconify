package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCheckCirclePrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><g opacity=".8"><path fill-rule="evenodd" d="M8 4.351c0-.47.414-.851.926-.851h6.148c.512 0 .926.381.926.851V7.65c0 .47-.414.851-.926.851H8.926C8.414 8.5 8 8.119 8 7.649V4.35Z" clip-rule="evenodd"/><path d="M6.462 19h10.576c.532 0 .962-.448.962-1V6c0-.552-.43-1-.962-1H6.462C5.93 5 5.5 5.448 5.5 6v12c0 .552.43 1 .962 1Z"/><path fill-rule="evenodd" d="M14.84 20c2.85 0 5.16-2.322 5.16-5.186S17.69 9.63 14.84 9.63s-5.16 2.321-5.16 5.185c0 2.864 2.31 5.186 5.16 5.186Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M6.18 2.5a.5.5 0 0 1 .5-.5h6.643a.5.5 0 0 1 .5.5v3.875a.5.5 0 0 1-.5.5H6.68a.5.5 0 0 1-.5-.5V2.5Zm1 .5v2.875h5.643V3H7.18Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3.305 4.507c0-.335.272-.607.607-.607h2.822v1.214H4.519v12.072h7.13c.32.49.745.905 1.243 1.214h-8.98a.607.607 0 0 1-.607-.607V4.507Zm13.393 6.87v-6.87a.607.607 0 0 0-.607-.607h-2.822v1.214h2.215v5.915c.429.052.837.171 1.214.348Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15 18a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 1a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M16.582 13.882a.5.5 0 0 1 .078.703l-1.106 1.382a1 1 0 0 1-1.488.082l-.696-.695a.5.5 0 0 1 .708-.707l.696.695l1.105-1.382a.5.5 0 0 1 .703-.078Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}