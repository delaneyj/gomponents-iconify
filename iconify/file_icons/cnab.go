package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cnab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M244.27 287.051h165.606v165.607H244.27V287.05zM414.016 59.342H269.111v111.785L202.868 38.64L99.364 245.65h207.008l-20.7-41.402h128.344V59.342zm-262.21 0L59.342 238.75V59.342h92.464zM0 0v144.906h20.7V20.7h124.206V0H0zm324.313 0v20.7h124.205v124.206h20.7V0H324.313zm124.205 367.094V491.3H324.313V512h144.905V367.094h-20.7zM197.25 400.966c39.647-68.672-10.29-155.04-89.672-155.09s-129.21 86.254-89.476 154.976s139.501 68.785 179.149.113z"/>`),
		g.Group(children),
	)
}