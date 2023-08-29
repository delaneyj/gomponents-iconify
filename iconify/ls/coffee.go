package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M397 344c-3-67-7-103-47-147c-51-57-70-112-28-197c11 57 15 108 55 139s73 115 20 205zm-75 0c-2-36-4-54-26-79c-27-30-38-59-15-105c6 31 7 58 29 74s41 62 12 110zm181 316h182c11 0 15 8 8 17c0 0-71 91-169 91H173L4 677c-7-9-4-17 7-17h180c-25-13-62-70-62-141V369h438v12c14-12 31-18 49-18c25 0 72 26 72 71c0 65-43 89-68 106c0 0-47 37-80 80c-11 19-24 34-37 40zm113-270c-16 0-36 7-49 39v83c0 11-1 24-3 35c22-22 48-37 48-37c16-11 48-32 48-73c1-29-26-47-44-47zm-411 12h-28v41h28v-41zm0 65h-28v67s8 90 65 117c-35-41-37-88-37-117v-67z"/>`),
		g.Group(children),
	)
}