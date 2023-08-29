package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiWhiteFlag0" d="M23.084 17.488C25.97 15.69 31.506 13.255 37.5 16c3.689 1.69 5.634 4.698 15.04 4.114c.79-.049 1.46.597 1.46 1.41V36.65c0 .623-.423 1.077-.965 1.354C51.088 39 45.764 38.94 37.5 35c-3.365-1.604-8.57-2.568-14.458 1.442"/></defs><g fill="#FFF"><path d="M22.302 17.488c2.958-1.798 8.635-4.233 14.78-1.488c3.783 1.69 5.778 4.698 15.422 4.114A1.414 1.414 0 0 1 54 21.524V36.65c0 .623-.434 1.077-.99 1.354c-1.995.996-7.454.935-15.927-3.004c-3.45-1.604-8.787-2.568-14.824 1.442"/><use href="#openmojiWhiteFlag0"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2"><use href="#openmojiWhiteFlag0" stroke-linejoin="round"/><path d="M19.327 12.875v46.96"/></g>`),
		g.Group(children),
	)
}