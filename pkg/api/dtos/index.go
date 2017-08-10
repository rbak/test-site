package dtos

type IndexViewData struct {
    AppUrl                  string
    AppSubUrl               string
    MainNavLinks            []*NavLink
}

type NavLink struct {
    Text     string     `json:"text,omitempty"`
    Icon     string     `json:"icon,omitempty"`
    Img      string     `json:"img,omitempty"`
    Url      string     `json:"url,omitempty"`
    Divider  bool       `json:"divider,omitempty"`
    Children []*NavLink `json:"children,omitempty"`
}
