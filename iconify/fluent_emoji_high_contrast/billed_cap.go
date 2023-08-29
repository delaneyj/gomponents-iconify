package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BilledCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.96 4.02v-.63c0-.77-.62-1.39-1.39-1.39h-1.18C14.62 2 14 2.62 14 3.39v.63h-2.93a7.75 7.75 0 0 0-7.75 7.75v5.93a3.954 3.954 0 0 0 2.682 3.742c.041.018.185.075.456.124c.262.055.534.084.812.084h17.492a3.927 3.927 0 0 0 1.158-.189a3.954 3.954 0 0 0 2.74-3.761v-5.93a7.75 7.75 0 0 0-7.75-7.75h-2.95Zm-1.95 17.63v-1H7.516a5.99 5.99 0 0 1-.527-.021V10.93a5.905 5.905 0 0 1 5.91-5.91h6.11a5.905 5.905 0 0 1 5.91 5.91v9.714a1.989 1.989 0 0 1-.158.006h-8.75v1Zm8.7 1H7.27a4.951 4.951 0 0 1-4.393-2.668L2.06 24.92c-.44 2.65 1.6 5.07 4.29 5.07h19.28c2.7 0 4.75-2.44 4.28-5.1l-.802-4.918a4.952 4.952 0 0 1-4.398 2.678Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}