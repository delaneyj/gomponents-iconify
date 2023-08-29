package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spotify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M88 408q59-85 162-85q92 0 151 71q67-67 67-162T400.5 69.5T237 2Q142 2 74.5 69.5T7 232q0 105 81 176zm273-107q-5 11-17 6q-47-22-101-22q-55 0-103 23q-10 6-14-7q-3-11 5-17q53-25 112-25q62 0 111 25q12 5 7 17zm19-59q-3 6-10 6q-3 0-4-1q-60-26-123-26q-64 0-122 26l-1 1h-3q-5 0-11-6l-2-8q0-6 5-9q63-29 134-29q72 0 135 29q5 3 5 9zM92 149q72-29 151-29t151 29q10 5 10 17q0 8-5.5 13.5T385 185q-3 0-4-1q-66-26-138-26q-73 0-137 26q-1 1-4 1q-8 0-14-5.5T82 166q0-10 10-17zm158 250q-60 0-96 48q39 15 83 15q54 0 102-23q-35-40-89-40z"/>`),
		g.Group(children),
	)
}