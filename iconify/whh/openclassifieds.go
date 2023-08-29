package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openclassifieds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 896q-149 0-269-91q-39 58-84.5 116T160 996q-28 27-66.5 27t-66-27.5t-27.5-66T27 864q17-17 75.5-62.5T218 717q-90-120-90-269q0-91 35.5-174T259 131t143-95.5T576 0t174 35.5T893 131t95.5 143t35.5 174t-35.5 174T893 765t-143 95.5T576 896zm0-768q-87 0-160.5 43T299 287.5T256 448t43 160.5T415.5 725T576 768t160.5-43T853 608.5T896 448t-43-160.5T736.5 171T576 128zm0 448q53 0 91-38l90 91q-75 75-181 75t-181-75t-75-181t75-181t181-75t181 75l-90 90q-38-37-91-37t-90.5 37.5T448 448t37.5 90.5T576 576z"/>`),
		g.Group(children),
	)
}