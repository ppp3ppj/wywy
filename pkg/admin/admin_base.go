package admin

import pub_variables "github.com/ppp3ppj/wywy/pub/variables"

type AuthNeed string

func AppendDeaderButton(userId uint) []pub_variables.InlineButton{
    baseUrl := "localhost"
    return []pub_variables.InlineButton{
        {
            AnchorUrl: baseUrl + "/",
            UserId: userId,
            IsBoosted: true,
            BaseUrl: baseUrl,
        },
    }
}
