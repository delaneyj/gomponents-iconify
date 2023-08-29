package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoldersyncPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.506 30.123c2.763.019 4.13-.151 5.068-1.298c1.364-1.667 1.047-3.89.168-4.862c-.88-.972-1.539-1.128-1.539-1.128s.163-1.11.163-1.848c.037-2.904-1.795-5.2-4.13-6.35c-2.057-.988-5.417-.618-6.923.958l-.61.638s-1.284-.788-1.856-1.06c-2.62-1.235-5.622-.783-7.563 1.288c-.282.31-.767 1.123-.767 1.123s-.73-.217-1.16-.175c-2.274.218-3.68 1.652-3.659 3.94c.005.458.186 1.37.186 1.37s-.556.168-.935.416c-1.068.7-1.549 1.868-1.432 3.479c.102 1.405.512 2.165 1.614 2.989c.524.391 1.018.611 3.19.567c2.503-.051 2.826-.081 2.826-.081s-.35 1.301 1.095 2.316c1.32.927 4.192 1.537 5.741 1.504c1.549-.033 3.977-.714 5.238-1.734c1.26-1.02 2.105-2.154 2.105-2.154s1.795.093 3.18.102Z"/>`),
		g.Group(children),
	)
}