package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerMediumVolume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M47.461 31.998c0-5.486-3.18-10.245-7.794-12.53L37.039 22.1c4.058 1.459 6.967 5.345 6.967 9.898c0 4.537-2.886 8.412-6.919 9.886l2.624 2.624c4.59-2.294 7.75-7.038 7.75-12.51"/><path fill="currentColor" d="M39.5 32c0-3.391-2.938-6.138-6.561-6.138V3.388c0-.767-.665-1.388-1.485-1.388c-.818 0-1.483.621-1.483 1.388v.697c-1.083.879-17.54 15.728-17.54 15.728H5.545C3.587 19.813 2 21.296 2 23.128v17.741c0 1.833 1.587 3.316 3.545 3.316h6.887v.002s9.631 7.364 17.54 15.705v.718c0 .767.665 1.39 1.483 1.39c.82 0 1.485-.623 1.485-1.39V38.136c3.623 0 6.56-2.748 6.56-6.136"/>`),
		g.Group(children),
	)
}