package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#F2F2F2" fill-rule="evenodd" d="M75 87.425L50 100L25 87.425L0 100V12.576L25 0l25 12.576L75 0l25 12.576V100L75 87.425z" clip-rule="evenodd"/><path fill="none" stroke="#6BC9F2" stroke-miterlimit="10" stroke-width="4" d="M15 60V32l9.988-5.006L50 41l25-12l12 4" clip-rule="evenodd"/><path fill="none" stroke="#E64C3C" stroke-miterlimit="10" stroke-width="4" d="M15 61v-8l10-5l25 13l25-11l12-5V32" clip-rule="evenodd"/><path fill="none" stroke="#F29C1F" stroke-miterlimit="10" stroke-width="4" d="m15 61l35 18l17-8V43l20-10" clip-rule="evenodd"/><path fill="#fff" fill-rule="evenodd" d="M87 36.5c-1.93 0-3.5-1.57-3.5-3.5s1.57-3.5 3.5-3.5s3.5 1.57 3.5 3.5s-1.57 3.5-3.5 3.5z" clip-rule="evenodd"/><path fill="#2980BA" d="M87 31c1.103 0 2 .897 2 2s-.897 2-2 2s-2-.897-2-2s.897-2 2-2m0-3a5 5 0 1 0 .001 10.001A5 5 0 0 0 87 28z"/><path fill="#fff" fill-rule="evenodd" d="M15 64.5c-1.93 0-3.5-1.57-3.5-3.5s1.57-3.5 3.5-3.5s3.5 1.57 3.5 3.5s-1.57 3.5-3.5 3.5z" clip-rule="evenodd"/><path fill="#2980BA" d="M15 59c1.103 0 2 .897 2 2s-.897 2-2 2s-2-.897-2-2s.897-2 2-2m0-3a5 5 0 1 0 .001 10.001A5 5 0 0 0 15 56z"/><path fill="#2C3E50" fill-rule="evenodd" d="m0 100l25-12.576V0L0 12.576V100zm50-87.424V100l25-12.576V0L50 12.576z" clip-rule="evenodd" opacity=".15"/>`),
		g.Group(children),
	)
}