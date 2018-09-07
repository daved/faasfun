♯卿fn

⋍卿∎
芽"encoding/json"
芽"fmt"
芽"net/http"
芽"time"
∘

⊯卿request卿⋥卿勇
芽Name卿string卿`json⟚"name"`
再

⊯卿response卿⋥卿勇
芽Msg卿卿string卿`json⟚"msg"`
芽Time卿string卿`json⟚"time"`
再

//卿Handler卿prints卿a卿hello卿message≛
⊲卿Handler∎w卿http≛ResponseWriter⦁卿r卿⊠http≛Request∘卿勇
芽d卿⟚⥀卿json≛NewDecoder∎r≛Body∘
芽defer卿r≛Body≛Close∎∘卿//卿nolint

芽req卿⟚⥀卿request勇再
芽if卿err卿⟚⥀卿d≛Decode∎⊕req∘;卿err卿!⥀卿nil卿勇
芽芽s卿⟚⥀卿http≛StatusInternalServerError
芽芽if卿_⦁卿ok卿⟚⥀卿err≛∎⊠json≛SyntaxError∘;卿ok卿勇
芽芽芽s卿⥀卿http≛StatusBadRequest
芽芽再
芽芽http≛Error∎w⦁卿http≛StatusText∎s∘⦁卿s∘
芽芽⊿
芽再

芽msg卿⟚⥀卿fmt≛Sprintf∎"Hello卿from卿GCP⦁卿%s≛卿Have卿a卿wonderful卿day≛"⦁卿req≛Name∘
芽resp卿⟚⥀卿response勇
芽芽Msg⟚卿卿msg⦁
芽芽Time⟚卿time≛Now∎∘≛String∎∘⦁
芽再
芽if卿err卿⟚⥀卿json≛NewEncoder∎w∘≛Encode∎resp∘;卿err卿!⥀卿nil卿勇
芽芽s卿⟚⥀卿http≛StatusInternalServerError
芽芽http≛Error∎w⦁卿http≛StatusText∎s∘⦁卿s∘
芽芽⊿
芽再
再
