/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { ComunService } from './comun.service';

describe('Service: Comun', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ComunService]
    });
  });

  it('should ...', inject([ComunService], (service: ComunService) => {
    expect(service).toBeTruthy();
  }));
});
