package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zodiaccancer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M544 960q32-64 32-160q92-53 161.5-129.5T847 499q-39 13-79 13q-106 0-181-75t-75-181t75-181T768 0t181 75t75 181q0 100-28.5 198.5T919 634T809.5 783t-130 113T544 960zm224-832q-53 0-90.5 37.5T640 256t37.5 90.5T768 384t90.5-37.5T896 256t-37.5-90.5T768 128zM177 525q38-13 79-13q106 0 181 75t75 181t-75 181t-181 75t-181-75T0 768q0-100 28.5-198.5T105 390t109.5-149t130-113T480 64q-32 64-32 160q-92 53-161.5 129.5T177 525zm-49 243q0 53 37.5 90.5T256 896t90.5-37.5T384 768t-37.5-90.5T256 640t-90.5 37.5T128 768z"/>`),
		g.Group(children),
	)
}