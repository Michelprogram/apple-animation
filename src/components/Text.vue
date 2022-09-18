<script setup lang="ts"></script>

<template>
  <div class="container-text">
    <div>
        <h2 class="title">Superpuissance <br>pour les pros.</h2>
        <p class="text">
        Le plus puissant des MacBook Pro est arrivé. Avec la puce M1 Pro ou M1
        Max, la première conçue par Apple pour les pros, vous bénéficiez de
        performances exceptionnelles et d’une incroyable autonomie. Ajoutez à cela
        un sublime écran Liquid Retina XDR, la meilleure caméra et le meilleur
        système audio jamais intégrés à un portable Mac et tous les ports dont
        vous avez besoin. Premier portable du genre, le MacBook Pro est un
        véritable monstre.
        </p>
    </div>
    <div class="price">
        <p>À partir de 2 249 €</p>
    </div>
    <div class="informations">
        <div>
            <p>Regarder l'évènement</p>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512">
                <path fill="#0071e3" d="M342.6 233.4c12.5 12.5 12.5 32.8 0 45.3l-192 192c-12.5 12.5-32.8 12.5-45.3 0s-12.5-32.8 0-45.3L274.7 256 105.4 86.6c-12.5-12.5-12.5-32.8 0-45.3s32.8-12.5 45.3 0l192 192z"/>
            </svg>
        </div>
        <div>
            <p>Regarder le film</p>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
               <path fill="#0071e3" d="M188.3 147.1C195.8 142.8 205.1 142.1 212.5 147.5L356.5 235.5C363.6 239.9 368 247.6 368 256C368 264.4 363.6 272.1 356.5 276.5L212.5 364.5C205.1 369 195.8 369.2 188.3 364.9C180.7 360.7 176 352.7 176 344V167.1C176 159.3 180.7 151.3 188.3 147.1V147.1zM512 256C512 397.4 397.4 512 256 512C114.6 512 0 397.4 0 256C0 114.6 114.6 0 256 0C397.4 0 512 114.6 512 256zM256 48C141.1 48 48 141.1 48 256C48 370.9 141.1 464 256 464C370.9 464 464 370.9 464 256C464 141.1 370.9 48 256 48z"/>
            </svg>
        </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>

@use "../assets/index.scss" as color;

.container-text{
    z-index:1;
    color: color.$title-color;

    &>div:nth-child(1){
        margin-bottom: 2%;;

        .title,.text{
            font-weight: 600;
            text-align: center;
        }

        .title{
            font-size: 4.5em;
            margin-bottom:1.5%;
        }

        .text{
            font-size: 1.5em;
            margin-left: 25%;
            margin-right: 25%;
            line-height: 1.125;
            letter-spacing:.004em;
        }
    }

    .price{
        margin-bottom:2%;
        p{
            font-size: 1.2em;
            color: color.$title-color;
            text-align: center;
        }
    }

    .informations{
        font-size:1.1em;
        width: 30%;
        height:50px;
        margin: auto;
        display: flex;
        justify-content: space-around;

        color: color.$header-blue-button;
        font-weight: 600;

        div{
            display: flex;
            align-items: center;
            cursor: pointer;

            p{
                white-space: nowrap;
                &:hover{
                    text-decoration: underline;
                }
            }
        }

        svg{
            width: 20px;
            margin-left: 2%;
        }


    }
}


</style>
