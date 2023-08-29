package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 704v64q0 106-75 181t-181 75H256q-106 0-181-75T0 768v-64q0-56 23-106t64-86q-41-36-64-86T0 320v-64Q0 150 75 75T256 0h256q106 0 181 75t75 181v64q0 56-23.5 106T681 512q40 36 63.5 86T768 704zM640 256q0-53-37.5-90.5T512 128H256q-53 0-90.5 37.5T128 256v64q0 53 37.5 90.5T256 448h256q53 0 90.5-37.5T640 320v-64zm0 448q0-53-37.5-90.5T512 576H256q-53 0-90.5 37.5T128 704v64q0 53 37.5 90.5T256 896h256q53 0 90.5-37.5T640 768v-64z"/>`),
		g.Group(children),
	)
}