package renderers

type RoutesMapType struct {
	Root       string
	Pinned     string
	Categories string
}

var RoutesMap RoutesMapType = RoutesMapType{
	Root:       "/",
	Pinned:     "/pinned/",
	Categories: "/categories/",
}
