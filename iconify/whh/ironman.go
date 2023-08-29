package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ironman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M773 832q-24 52-73 102t-109 90H241q-56-37-103.5-88T64 832Q0 704 0 320q0-35 11-96.5T32 128l160-96Q291 0 419 0q121 0 221 32l160 96q10 34 21 95.5t11 96.5q0 384-59 512zM271 960h294q88-61 139-160l-64 32h-64l-48-64H304l-48 64h-64l-64-32q43 90 143 160zm497-768L606 92q-33-11-65-17l-29 117q-7 27-23 45.5T448 256h-64q-24 0-40.5-19.5T320 192L291 76q-30 6-60 16L64 192v224q0 55 16 137.5T96 672q0 24 6 51l125 62l61-81h256l61 81l125-62q6-28 6-51q0-34 16-118.5T768 416V192zM512 416l192-32v64l-192 64v-96zm-384-32l192 32v96l-192-64v-64z"/>`),
		g.Group(children),
	)
}