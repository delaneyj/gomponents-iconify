package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3h.01M7 3h.01M17 3h.01"/><path fill="currentColor" fill-rule="evenodd" d="M7 5a1 1 0 0 1 1 1v2h3V6a1 1 0 1 1 2 0v2h3V6a1 1 0 1 1 2 0v2.1a5.005 5.005 0 0 1 3.937 4.102c-.4.155-.75.383-1.047.63c-.532.44-.966.994-1.302 1.46c-.265.367-.714.708-1.588.708c-.874 0-1.324-.342-1.588-.71A2.355 2.355 0 0 1 16 13a1 1 0 1 0-2 0c0 .34-.11.872-.412 1.29c-.264.368-.714.71-1.588.71c-.874 0-1.324-.342-1.588-.71A2.355 2.355 0 0 1 10 13a1 1 0 1 0-2 0c0 .34-.11.872-.412 1.29c-.264.368-.714.71-1.588.71c-.874 0-1.323-.341-1.588-.709c-.336-.465-.77-1.019-1.302-1.46a3.822 3.822 0 0 0-1.047-.629A5.005 5.005 0 0 1 6 8.1V6a1 1 0 0 1 1-1zm-5 9.52V19a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4.48c-.27.256-.532.583-.79.94c-.635.883-1.685 1.54-3.21 1.54c-1.367 0-2.353-.529-3-1.273c-.647.744-1.633 1.273-3 1.273s-2.353-.529-3-1.273C8.353 16.47 7.367 17 6 17c-1.525 0-2.575-.657-3.21-1.54a6.756 6.756 0 0 0-.79-.94z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}