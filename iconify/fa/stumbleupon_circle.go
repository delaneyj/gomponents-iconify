package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StumbleuponCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m866 711l90-27v-62q0-79-58-135t-138-56t-138 55.5T564 621v283q0 20-14 33.5T517 951t-32.5-13.5T471 904V784H320v122q0 82 57.5 139t139.5 57q81 0 138.5-56.5T713 909V629q0-19 13.5-33t33.5-14q19 0 32.5 14t13.5 33v54zm333 195V784h-150v126q0 20-13.5 33.5T1002 957q-19 0-32.5-14T956 910V787l-90 26l-60-28v123q0 80 58 137t139 57t138.5-57t57.5-139zm337-138q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768z"/>`),
		g.Group(children),
	)
}