package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterFrames(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M384 104H128v236.117h256Zm-32 204.117H160V136h192Z"/><path fill="currentColor" d="M181.49 428.117L256 502.628l74.51-74.511H440a32.036 32.036 0 0 0 32-32V48a32.036 32.036 0 0 0-32-32H72a32.036 32.036 0 0 0-32 32v348.117a32.036 32.036 0 0 0 32 32ZM72 48h368v348.117H317.255L256 457.372l-61.255-61.255H72Z"/>`),
		g.Group(children),
	)
}