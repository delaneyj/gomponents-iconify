package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraRetro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M928 704q0-14-9-23t-23-9q-66 0-113 47t-47 113q0 14 9 23t23 9t23-9t9-23q0-40 28-68t68-28q14 0 23-9t9-23zm224 130q0 106-75 181t-181 75t-181-75t-75-181t75-181t181-75t181 75t75 181zM128 1408h1536v-128H128v128zm1152-574q0-159-112.5-271.5T896 450T624.5 562.5T512 834t112.5 271.5T896 1218t271.5-112.5T1280 834zM256 192h384V64H256v128zM128 384h1536V128H836l-64 128H128v128zm1664-256v1280q0 53-37.5 90.5T1664 1536H128q-53 0-90.5-37.5T0 1408V128q0-53 37.5-90.5T128 0h1536q53 0 90.5 37.5T1792 128z"/>`),
		g.Group(children),
	)
}