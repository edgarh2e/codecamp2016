import {Component, OnInit} from '@angular/core';
import {ComunService} from "./service/comun.service";
import {Usuario} from "./domain/usuario";
import {Respuesta} from "./domain/respuesta";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  usuario1: string;
  usuario2: string;
  habilitaGrafo: boolean = false;
  usuarios: Usuario[] = [];
  respuesta: Respuesta;
  rutaSVG: string;
  imagenurl: string = 'http://45.55.217.15:3000/compare/view/ce18eeaee1d1a1d02a49bc8201a817f7.svg';

  constructor(private comunService: ComunService){}

  ngOnInit() {

  }


  generarGrafo(valor: any) {
    console.log('Generar grafo,',valor);
    this.usuarios[0] = new Usuario('',valor.usu1);
    this.usuarios[1] = new Usuario('',valor.usu2);
    this.comunService.obtenerSVG(this.usuarios).subscribe(
      (respuesta: Respuesta) => {
        this.respuesta = respuesta;
        this.rutaSVG = 'http://45.55.217.15:3000'+respuesta.image;
        console.log(this.rutaSVG);
      }
    );
    this.habilitaGrafo = true;
  }

}
