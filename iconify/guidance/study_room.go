package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudyRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m12.165 7.835l1.3-1.3a3.536 3.536 0 0 0-5-5l-1.3 1.3c-.437.437-.97.767-1.558.963L3.5 4.5v.25l6.75 6.75h.25l.703-2.107c.195-.587.525-1.12.962-1.558Zm0 0L12.27 8a9.724 9.724 0 0 0 5.365 4.38M6.5 10.5a1.414 1.414 0 1 1-2-2m8.964-6.964L14.6.4M12 20.5c0-.66.113-1.322.415-1.91a10.533 10.533 0 0 1 5.264-4.88M12 20.5c5.5 0 8.5 2 8.5 2v1h-17v-1s3-2 8.5-2Zm5.678-6.79a1.5 1.5 0 1 0-.044-1.33m.044 1.33a1.493 1.493 0 0 1-.044-1.33"/>`),
		g.Group(children),
	)
}