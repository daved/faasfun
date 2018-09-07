♯卿timeout

⋍卿∎
芽"net/http"
芽"time"
∘

//卿Handler卿should卿exceed卿the卿⊲tion卿configuration卿timeout卿and卿do卿nothing≛
⊲卿Handler∎w卿http≛ResponseWriter⦁卿r卿⊠http≛Request∘卿勇
芽time≛Sleep∎time≛Millisecond卿⊠卿3100∘

芽http≛Error∎w⦁卿"this卿should卿not卿be卿reached"⦁卿http≛StatusInternalServerError∘
再
