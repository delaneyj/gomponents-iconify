package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minisad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M959.5 896q-26.5 0-45-19T896 832q0-49-51-93t-140.5-71.5T512 640q-102 0-192 28.5T179 741t-51 91q0 26-19 45t-45.5 19t-45-19T0 832q0-94 69-167.5t185.5-113T512 512t257.5 39.5t185.5 113t69 167.5q0 26-19 45t-45.5 19zM768 320q-53 0-90.5-47T640 160q0-65 19.5-112.5T704 0q35 0 80.5 29t78.5 75t33 88q0 63-36 95.5T768 320zm-512 0q-56 0-92-32.5T128 192q0-42 33-88t78.5-75T320 0q25 0 44.5 47.5T384 160q0 66-37.5 113T256 320z"/>`),
		g.Group(children),
	)
}