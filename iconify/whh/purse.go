package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Purse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 960q-46 23-98 37t-117 19.5t-109.5 6.5t-123.5 1q-81 0-123-1.5t-108.5-6.5T162 997t-98-37q-31-16-47.5-53T0 832q0-53 41.5-205.5T128 384q12-25 28-41t26-19l10-4q22 0 64 11v-11q0-87 34.5-160.5t93-116.5T512 0t129 43t93 116.5T768 320v11q43-11 64-11q19 0 33 15t31 49q45 90 86.5 242.5T1024 832q1 96-64 128zM320 448q-27 0-45.5 18.5t-18.5 45t18.5 45.5t45.5 19t45.5-19t18.5-45.5t-18.5-45T320 448zM512 63q-53 0-90.5 70.5T384 304q0 29 5 64q70 16 123 16t123-16q5-35 5-64q0-100-37.5-170.5T512 63zm192 385q-27 0-45.5 18.5t-18.5 45t18.5 45.5t45.5 19t45.5-19t18.5-45.5t-18.5-45T704 448z"/>`),
		g.Group(children),
	)
}