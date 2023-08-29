package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChargingStation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 5C7.355 5 6 6.355 6 8v19h14v-7h2v4c0 1.645 1.355 3 3 3s3-1.355 3-3V13.812c0-.796-.313-1.531-.875-2.093L22.406 7L21 8.406l2.844 2.844C22.77 11.707 22 12.766 22 14c0 1.645 1.355 3 3 3a2.93 2.93 0 0 0 1-.188V24c0 .566-.434 1-1 1c-.566 0-1-.434-1-1v-4c0-1.094-.906-2-2-2h-2V8c0-1.645-1.355-3-3-3H9zm0 2h8c.566 0 1 .434 1 1v17H8V8c0-.566.434-1 1-1zm4.09 3.56l-2.43 4.88l.9.47l1.1.53l-1.57 3.12l1.82.88l2.43-4.88l-.9-.47l-1.1-.53l1.57-3.12l-1.82-.88zM25 13c.562 0 1 .438 1 1s-.438 1-1 1c-.563 0-1-.438-1-1s.437-1 1-1z"/>`),
		g.Group(children),
	)
}