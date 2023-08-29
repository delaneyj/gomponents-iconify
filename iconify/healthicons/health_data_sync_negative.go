package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthDataSyncNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsHealthDataSyncNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM20 4a3 3 0 0 0-3 3h-3a3 3 0 0 0-3 3v28a3 3 0 0 0 3 3h10v-2H14a1 1 0 0 1-1-1V10a1 1 0 0 1 1-1h3a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3h3a1 1 0 0 1 1 1v16h2V10a3 3 0 0 0-3-3h-3a3 3 0 0 0-3-3h-8Zm-1 3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1V7Zm4 12v-3h2v3h3v2h-3v3h-2v-3h-3v-2h3Zm1 14.665v-4h2v1.688l1.398-1.313a7 7 0 0 1 11.538 2.625l-1.885.667a5.001 5.001 0 0 0-8.273-1.845l-1.253 1.178H29v2h-4a1 1 0 0 1-1-1Zm15.667 1.667h-4v2h1.475l-1.254 1.178a5 5 0 0 1-8.273-1.845l-1.884.667a7 7 0 0 0 11.538 2.625l1.398-1.313v1.688h2v-4a1 1 0 0 0-1-1Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsHealthDataSyncNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}