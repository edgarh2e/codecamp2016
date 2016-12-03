import { Injectable } from '@angular/core';
import {Http, Response} from '@angular/http';
import {Usuario} from '../domain/usuario';
import {Respuesta} from '../domain/respuesta';
import {Observable} from 'rxjs';

@Injectable()
export class ComunService {

  constructor(private http: Http) { }

  obtenerSVG(usuarios: Usuario[]): Observable<Respuesta> {
    let ruta = 'http://45.55.217.15:3000/compare?'+this.generarArgumentos(usuarios);
    console.log(ruta);
    return this.http.get(ruta)
      .map((respuesta: Response) => {
        console.log('JSON: ',respuesta.json());
        return respuesta.json();
      })
      .catch(this.manejadorDeError);
  }

  private generarArgumentos(usuarios: Usuario[]): string{
    let parametros: string = '';
    let tamano = usuarios.length;
    for(let usuario of usuarios){
      parametros = parametros+'user='+usuario.nombre+'&';
    }
    if(parametros.endsWith('&')){
      parametros = parametros.slice(0,parametros.length-1);
    }
    console.log(parametros);
    return parametros;
  }

  public manejadorDeError(error: any) {
    return Observable.throw(error);
  }

}
