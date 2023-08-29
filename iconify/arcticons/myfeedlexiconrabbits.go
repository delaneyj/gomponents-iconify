package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myfeedlexiconrabbits(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.117 18.612l-4.703-9.06a2.369 2.369 0 0 1 .207-2.005a3.958 3.958 0 0 1 4.634-2.006c.74.257 1.706.808 1.729 1.59c0 0-.164 11.025 0 11.066Zm4.218-.484s1.301-8.203 1.798-10.374a2.557 2.557 0 0 1 1.937-1.729a4.99 4.99 0 0 1 4.011 1.452a2.05 2.05 0 0 1 .484 2.42c-1.648 3.166-6.293 8.438-6.293 8.438Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.603 42.626h-22.96l.277-8.454s-.694-.129-1.037-.208c-1.695-.39-3.637-.299-5.049-1.314a9.69 9.69 0 0 1-3.388-5.186a5.907 5.907 0 0 1 .069-3.043a8.276 8.276 0 0 1 1.521-2.905a9.57 9.57 0 0 1 2.49-2.282a4.964 4.964 0 0 1 1.59-.622a26.594 26.594 0 0 1 6.156-.277a20.59 20.59 0 0 1 2.213.207a40.652 40.652 0 0 1 7.607 1.314a15.93 15.93 0 0 1 4.703 2.421a11.73 11.73 0 0 1 2.835 3.181a18.36 18.36 0 0 1 1.867 4.98a43.388 43.388 0 0 1 1.107 7.676c.107 1.5 0 4.512 0 4.512Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.625 41.894a6.504 6.504 0 0 0 1.984-.634a2.466 2.466 0 0 0 1.037-1.383a3.736 3.736 0 0 0-.691-3.043a5.883 5.883 0 0 0-2.555-1.086"/>`),
		g.Group(children),
	)
}