package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneSlashThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m211 213.31l-160-176a4 4 0 0 0-6 5.38l39 42.86V128a44 44 0 0 0 70.52 35.11l16.17 17.8A68 68 0 0 1 60 128a4 4 0 0 0-8 0a76.09 76.09 0 0 0 72 75.89V232a4 4 0 0 0 8 0v-28.11a75.76 75.76 0 0 0 44.09-17l28.91 31.8a4 4 0 1 0 5.92-5.38ZM128 164a36 36 0 0 1-36-36V94.35l57.1 62.81A35.9 35.9 0 0 1 128 164ZM90.56 40.88A44 44 0 0 1 172 64v60.43a4 4 0 0 1-8 0V64a36 36 0 0 0-66.64-18.91a4 4 0 1 1-6.8-4.21Zm100.81 111.85A67.81 67.81 0 0 0 196 128a4 4 0 0 1 8 0a75.64 75.64 0 0 1-5.18 27.64a4 4 0 0 1-3.73 2.54a3.88 3.88 0 0 1-1.45-.27a4 4 0 0 1-2.27-5.18Z"/>`),
		g.Group(children),
	)
}