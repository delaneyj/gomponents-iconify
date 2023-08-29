package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCheckCircledCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g opacity=".2"><path fill-rule="evenodd" d="M11 7.351c0-.47.414-.851.926-.851h6.148c.512 0 .926.381.926.851v3.298c0 .47-.414.851-.926.851h-6.148c-.512 0-.926-.381-.926-.851V7.35Z" clip-rule="evenodd"/><path d="M9.462 22h10.576c.532 0 .962-.448.962-1V9c0-.552-.43-1-.962-1H9.462C8.93 8 8.5 8.448 8.5 9v12c0 .552.43 1 .962 1Z"/><path d="M23 18.75a4.25 4.25 0 1 1-8.5 0a4.25 4.25 0 0 1 8.5 0Z"/></g><path fill-rule="evenodd" d="M9.175 5.5a.5.5 0 0 1 .5-.5h6.643a.5.5 0 0 1 .5.5v3.875a.5.5 0 0 1-.5.5H9.675a.5.5 0 0 1-.5-.5V5.5Zm1 .5v2.875h5.643V6h-5.643Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.5 20V8h2V7h-2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h7.854a4.02 4.02 0 0 1-.819-1H7.5Zm11-5.97c.35.045.685.133 1 .26V8a1 1 0 0 0-1-1h-2v1h2v6.03Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M18 21a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 1a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M19.582 16.882a.5.5 0 0 1 .078.703l-1.106 1.382a1 1 0 0 1-1.488.082l-.695-.695a.5.5 0 0 1 .707-.707l.695.695l1.106-1.382a.5.5 0 0 1 .703-.078Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}